package controllers

import (
	"net/http"
	"task-manager/config"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, description, status FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}
func CreateTask(c *gin.Context) {
	var task models.Task

	// Bind JSON request
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Insert into DB
	query := `INSERT INTO tasks (title, description, status) 
	          VALUES ($1, $2, $3) RETURNING id`

	err := config.DB.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Status,
	).Scan(&task.ID)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, task)
}
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task models.Task

	// Bind JSON
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	query := `
	UPDATE tasks 
	SET title=$1, description=$2, status=$3 
	WHERE id=$4
	`

	_, err := config.DB.Exec(
		query,
		task.Title,
		task.Description,
		task.Status,
		id,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task updated successfully"})
}
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	query := `DELETE FROM tasks WHERE id=$1`

	_, err := config.DB.Exec(query, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task deleted successfully"})
}