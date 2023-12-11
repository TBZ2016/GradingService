package grading_test

import (
	"fmt"
	"os"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

var opts = godog.Options{
	Format: "progress", // you can specify other formatters
	Paths:  []string{"."},
	Output: colors.Colored(os.Stdout),
}

func thereIsAGradingService() error {
	// setup grading service for testing
	return nil
}

func iRequestGradesForCursusID(cursusID int) error {
	// make a request to the grading service for grades by cursus ID
	return nil
}

func gradingServiceShouldReturnListOfGrades() error {
	// check if the grading service returned the expected list of grades
	return fmt.Errorf("not implemented")
}

func FeatureContext(s *godog.Suite) {
	s.BeforeScenario(func(interface{}) {
		// reset state before each scenario
	})

	s.Step(`^there is a grading service$`, thereIsAGradingService)
	s.Step(`^I request grades for cursus ID (\d+)$`, iRequestGradesForCursusID)
	s.Step(`^grading service should return a list of grades$`, gradingServiceShouldReturnListOfGrades)
}

func ExampleFeature() {
	opts.Paths = []string{"grading_service.feature"}
	godog.TestSuite{
		Name:                "features",
		ScenarioInitializer: FeatureContext,
		Options:             &opts,
	}.Run()
}
