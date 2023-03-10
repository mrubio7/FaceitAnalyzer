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
	maxIter := 1000

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
	theta := []float64{-0.08180087541815817, -0.054204510847769734, -0.2699163397317514, -0.07445649606263728, -0.051159913533702625, -0.26136168795783193, -0.07739696171776977, -0.05235787009725521, -0.25944417573122946, -0.09270077569475715, -0.0614981518312316, -0.356938547623408, -0.07476328621927235, -0.0508460436042222, -0.2604471435509979, -0.07578778834811205, -0.0507012032514262, -0.26965084825004815}
	newSample := data.Data
	prediction := predict(theta, newSample)

	return prediction
}
