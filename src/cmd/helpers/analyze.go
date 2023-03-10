package helpers

import (
	"faceitAI/src/models"
	"fmt"
	"math"
)

var Analyze iAnalyze = analyze{}

type iAnalyze interface {
	Training(data []models.Results)
	Predict(data models.Results) float64
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

func predict(theta []float64, features []float64) float64 {
	return sigmoid(dotProduct(theta, features))
}

func dotProduct(a []float64, b []float64) float64 {
	sum := 0.0
	for i := 0; i < len(a); i++ {
		sum += a[i] * b[i]
	}
	return sum
}

func (a analyze) Training(data []models.Results) {
	alpha := 0.1
	lambda := 0.01
	maxIter := 10000

	n := len(data[0].Data)
	theta := make([]float64, n)

	for i := 0; i < maxIter; i++ {
		J := costFunction(theta, data, lambda)
		theta = gradientDescent(theta, data, alpha, lambda)
		fmt.Println(theta)
		fmt.Println("Iteración:", i+1, "Coste:", J)
	}

}

func (a analyze) Predict(data models.Results) float64 {

	theta := []float64{1.102069451377616, -1.8871836511889366, -0.14209854361641328, -0.0697153440145562, -1.2534425192033358, 0.2209949148579937, -1.7567309912532303, -1.515994738412653, 0.33329065417812376, -1.9316215049438357, -2.801961090988334, 0.2624296009807843, -0.4054550776242352, -1.1961529504387949, -1.4018474465197746, 0.34482608494502187, -0.6142893631274927, 3.277209457405923}
	newSample := data.Data
	prediction := predict(theta, newSample)

	return prediction
}
