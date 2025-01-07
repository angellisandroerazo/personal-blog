package controllers

import (
	"angellisandroerazo/personal-blog/models"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func listPosts() ([]models.Posts, string) {
	file, err := os.ReadFile("./posts/posts.json")
	if err != nil {

		file, err := os.Create("./posts/posts.json")
		if err != nil {
			return nil, "Error creating file: " + err.Error()
		}

		defer file.Close()

	}

	posts := []models.Posts{}

	json.Unmarshal(file, &posts)

	return posts, "Success"
}

func savePosts(posts []models.Posts) string {
	file, err := os.Create("./posts/posts.json")
	if err != nil {
		return "Error creating file: " + err.Error()
	}

	defer file.Close()

	data, err := json.Marshal(posts)
	if err != nil {
		return "Error parsing posts: " + err.Error()
	}

	err = os.WriteFile("./posts/posts.json", data, 0644)
	if err != nil {
		return "Error write posts: " + err.Error()
	}

	return "Success"
}

func CreatePage(c *fiber.Ctx) error {
	return c.Render("create", fiber.Map{})
}

func CreatePost(c *fiber.Ctx) error {
	if c.Method() != "POST" {
		return c.Status(405).SendString("Method Not Allowed")
	}

	posts, _ := listPosts()

	post := models.Posts{
		ID:    len(posts) + 1,
		Title: c.FormValue("title"),
		Date:  time.Now().Format("2006-01-02 15:04"),
		Body:  c.FormValue("body"),
	}

	posts = append(posts, post)

	savePosts(posts)

	return c.Redirect("/admin")
}

func ViewPost(c *fiber.Ctx) error {
	posts, _ := listPosts()

	id, _ := strconv.Atoi(c.Params("id"))

	for _, post := range posts {
		if post.ID == id {
			return c.Render("post", fiber.Map{
				"Post": post,
			})
		}
	}

	return c.Status(404).SendString("Post not found")
}

func UpdatePage(c *fiber.Ctx) error {
	posts, _ := listPosts()

	id, _ := strconv.Atoi(c.Params("id"))

	for _, post := range posts {
		if post.ID == id {
			return c.Render("update", fiber.Map{
				"Post": post,
			})
		}
	}

	return c.Status(404).SendString("Post not found")
}

func UpdatePost(c *fiber.Ctx) error {
	if c.Method() != "POST" {
		return c.Status(405).SendString("Method Not Allowed")
	}

	posts, _ := listPosts()

	id, _ := strconv.Atoi(c.Params("id"))

	for i, post := range posts {
		if post.ID == id {
			posts[i] = models.Posts{
				ID:    id,
				Title: c.FormValue("title"),
				Date:  posts[i].Date,
				Body:  c.FormValue("body"),
			}
		}
	}

	fmt.Println(posts)

	savePosts(posts)

	return c.Redirect("/admin")
}

func DeletePost(c *fiber.Ctx) error {
	posts, _ := listPosts()

	id, _ := strconv.Atoi(c.Params("id"))

	updatePosts := []models.Posts{}

	for _, post := range posts {
		if post.ID != id {
			updatePosts = append(updatePosts, post)
		}
	}

	savePosts(updatePosts)

	return c.Redirect("/admin")
}
