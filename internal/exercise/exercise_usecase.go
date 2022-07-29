package exercise

import (
	"course/internal/domain"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ExerciseUsecase struct {
	db *gorm.DB
}

func NewExerciseUsecase(db *gorm.DB) *ExerciseUsecase {
	return &ExerciseUsecase{db}
}

func (exUsecase ExerciseUsecase) GetExerciseByID(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, map[string]interface{}{
			"message": "invalid exercise id",
		})
		return
	}
	var exercise domain.Exercise
	err = exUsecase.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(404, map[string]interface{}{
			"message": "exercise not found",
		})
		return
	}
	c.JSON(200, exercise)
}

func (exUseCase ExerciseUsecase) CreateExercise(c *gin.Context) {
	var exercise domain.Exercise
	err := c.ShouldBind(&exercise)
	if err != nil {
		c.JSON(404, gin.H{
			"messages": "invalid input",
		})
		return
	}
	if exercise.Title == "" {
		c.JSON(404, gin.H{
			"messages": "title kosong",
		})
		return
	}
	if exercise.Description == "" {
		c.JSON(404, gin.H{
			"messages": "description kosong",
		})
		return
	}

	err = exUseCase.db.Create(&exercise).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed create exercise",
		})
		return
	}
	c.JSON(201, gin.H{
		"id":          exercise.ID,
		"title":       exercise.Title,
		"description": exercise.Description,
	})
}

func (exUsecase ExerciseUsecase) CreateAnswer(c *gin.Context) {
	var answer domain.Answer
	exerciseId := c.Param("exerciseId")
	exerId, err := strconv.Atoi(exerciseId)
	userId := int(c.Request.Context().Value("user_id").(float64))
	qustionId := c.Param("questionId")
	questionId, err1 := strconv.Atoi(qustionId)
	if err != nil {
		c.JSON(400, map[string]interface{}{
			"message": "bad request",
		})
		return
	}

	if err1 != nil {
		c.JSON(400, map[string]interface{}{
			"message": "bad request",
		})
		return
	}

	c.ShouldBind(&answer)
	answer.ExerciseID = exerId
	answer.UserID = userId
	answer.QuestionID = questionId
	validate := validator.New()
	err = validate.Struct(answer)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "input salah",
		})
		return
	}
	err = exUsecase.db.Create(&answer).Error
	if err != nil {
		c.JSON(400, gin.H{
			"message": "failed create answer",
		})
		return
	}

	c.JSON(200, gin.H{
		"messages": "success",
	})

}

func (exUsecase ExerciseUsecase) CreateQuestion(c *gin.Context) {
	var question domain.Question
	exerciseId := c.Param("exerciseId")
	id, err := strconv.Atoi(exerciseId)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, map[string]interface{}{
			"message": "id exercise not found",
		})
		return
	}
	question.ExerciseID = id
	userId := int(c.Request.Context().Value("user_id").(float64))
	question.CreatorID = userId
	err = c.ShouldBind(&question)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid input",
		})
		return
	}
	if question.Body == "" {
		c.JSON(400, gin.H{
			"message": "field body empty",
		})
		return
	}

	if question.OptionA == "" {
		c.JSON(400, gin.H{
			"message": "field option_a empty",
		})
		return
	}
	if question.OptionB == "" {
		c.JSON(400, gin.H{
			"message": "field option_b empty",
		})
		return
	}
	if question.OptionC == "" {
		c.JSON(400, gin.H{
			"message": "field option_c empty",
		})
		return
	}
	if question.OptionD == "" {
		c.JSON(400, gin.H{
			"message": "field option_d empty",
		})
		return
	}
	if question.CorrectAnswer == "" {
		c.JSON(400, gin.H{
			"message": "field correct_answer empty",
		})
		return
	}
	if question.Score == 0 {
		c.JSON(400, gin.H{
			"message": "field score empty",
		})
		return
	}

	err = exUsecase.db.Create(&question).Error
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed create question",
		})
		fmt.Println(err)
		return
	}
	c.JSON(201, gin.H{
		"message": "success create question",
	})
}
func (exUsecase ExerciseUsecase) CalculateUserScore(c *gin.Context) {
	paramID := c.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(400, map[string]interface{}{
			"message": "invalid exercise id",
		})
		return
	}
	var exercise domain.Exercise
	err = exUsecase.db.Where("id = ?", id).Preload("Questions").Take(&exercise).Error
	if err != nil {
		c.JSON(404, map[string]interface{}{
			"message": "exercise not found",
		})
		return
	}

	userID := int(c.Request.Context().Value("user_id").(float64))
	var answers []domain.Answer
	err = exUsecase.db.Where("user_id = ?", userID).Find(&answers).Error
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"message": "error when find answers",
		})
		return
	}
	if len(answers) == 0 {
		c.JSON(200, map[string]interface{}{
			"score": 0,
		})
		return
	}

	mapQuestion := make(map[int]domain.Question)
	for _, question := range exercise.Questions {
		mapQuestion[question.ID] = question
	}

	var score int
	for _, answer := range answers {
		if strings.EqualFold(answer.Answer, mapQuestion[answer.QuestionID].CorrectAnswer) {
			score += mapQuestion[answer.QuestionID].Score
		}
	}
	c.JSON(200, map[string]interface{}{
		"score": score,
	})
}
