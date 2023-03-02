package helpers

import (
	"faceitAI/src/models"
	"fmt"
	"math"
)

var Analyze iAnalyze = analyze{}

type iAnalyze interface {
	Calculate(data []float32)
}

type analyze struct {
	iAnalyze
}

func sigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

func costFunction(theta []float64, samples []models.Results, lambda float64) float64 {
	m := len(samples)
	J := 0.0
	for i := 0; i < m; i++ {
		h := sigmoid(dotProduct(theta, samples[i].Data))
		J += float64(-samples[i].Label)*math.Log(h) - float64(1-samples[i].Label)*math.Log(1-h)
	}
	J = J/float64(m) + lambda*dotProduct(theta, theta)/(2*float64(m))
	return J
}

func gradientDescent(theta []float64, samples []models.Results, alpha float64, lambda float64) []float64 {
	m := len(samples)
	n := len(theta)
	grad := make([]float64, n)
	for i := 0; i < m; i++ {
		h := sigmoid(dotProduct(theta, samples[i].Data))
		for j := 0; j < n; j++ {
			grad[j] += (h - float64(samples[i].Label)) * samples[i].Data[j]
		}
	}
	for j := 0; j < n; j++ {
		grad[j] = grad[j]/float64(m) + lambda*theta[j]/float64(m)
		theta[j] = theta[j] - alpha*grad[j]
	}
	return theta
}

func predict(theta []float64, features []float64) int {
	h := sigmoid(dotProduct(theta, features))
	if h >= 0.5 {
		return 1
	} else {
		return 0
	}
}

func dotProduct(a []float64, b []float64) float64 {
	sum := 0.0
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}

func Trainning(data []models.Results) {
	alpha := 0.1
	lambda := 0.01
	maxIter := 1000

	n := len(data[0].Data)
	theta := make([]float64, n)

	for i := 0; i < maxIter; i++ {
		J := costFunction(theta, data, lambda)
		theta = gradientDescent(theta, data, alpha, lambda)
		fmt.Println("Iteración:", i+1, "Coste:", J)
	}

}

func Predict(data []models.Results) {
	n := len(data[0].Data)
	theta := make([]float64, n)

	newSample := data[1].Data
	prediction := predict(theta, newSample)
	fmt.Println(prediction)
}
