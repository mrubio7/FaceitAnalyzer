package helpers

import (
	"faceitAI/src/models"
	"fmt"
	"math"
)

var Analyze iAnalyze = analyze{}

type iAnalyze interface {
	Training(data []models.Results)
	Predict(data models.Results)
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
	maxIter := 5000

	n := len(data[0].Data)
	theta := make([]float64, n)

	for i := 0; i < maxIter; i++ {
		J := costFunction(theta, data, lambda)
		theta = gradientDescent(theta, data, alpha, lambda)
		fmt.Println(theta)
		fmt.Println("Iteración:", i+1, "Coste:", J)
	}

}

func (a analyze) Predict(data models.Results) {

	theta := []float64{2.621290464066871, 0.4906159278448144, 1.271165474863099, 0.02594410857806587, 0.19981465381570995, -2.505114758295816, -0.7492416802685565, -0.172501877591028, -2.8772545207431603, -0.4035303506674547, -0.44031867535800534, -1.1223985283530968, -0.5162679352623092, -0.3250840593052962, 2.405242809372621, 0.8621031408213191, 0.15635298402142783, 2.225612434445203}

	newSample := data.Data
	prediction := predict(theta, newSample)
	fmt.Println(prediction)
}
