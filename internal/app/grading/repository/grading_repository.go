package repository

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	//"go.mongodb.org/mongo-driver/mongo/options"

	"kawa/gradingservice/internal/app/grading/model"
)

type GradingRepository struct {
	collection *mongo.Collection
}

func NewGradingMongoDBRepository(db *mongo.Database) *GradingRepository {
	return &GradingRepository{
		collection: db.Collection("grades"),
	}
}

func (r *GradingRepository) GetByCursusID(cursusID int) ([]model.Grade, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"courseId": cursusID}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var grades []model.Grade
	if err := cursor.All(ctx, &grades); err != nil {
		return nil, err
	}

	return grades, nil
}

func (r *GradingRepository) Create(grade *model.Grade) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, grade)
	if err != nil {
		return err
	}

	return nil
}

func (r *GradingRepository) GetByStudentID(studentID int) ([]model.Grade, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"studentId": studentID}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var grades []model.Grade
	if err := cursor.All(ctx, &grades); err != nil {
		return nil, err
	}

	return grades, nil
}

func (r *GradingRepository) GetByClass(classID int) ([]model.Grade, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Assuming you have a class field in your Grade struct
	filter := bson.M{"class": classID}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var grades []model.Grade
	if err := cursor.All(ctx, &grades); err != nil {
		return nil, err
	}

	return grades, nil
}

func (r *GradingRepository) GetById(gradeID int) (*model.Grade, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"gradeId": gradeID}

	var grade model.Grade
	err := r.collection.FindOne(ctx, filter).Decode(&grade)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // Return nil if not found
		}
		return nil, err
	}

	return &grade, nil
}

func (r *GradingRepository) Update(grade *model.Grade) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"gradeId": grade.GradeID}
	update := bson.M{"$set": grade}

	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *GradingRepository) DeleteById(gradeID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"gradeId": gradeID}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("grade not found")
	}

	return nil
}
