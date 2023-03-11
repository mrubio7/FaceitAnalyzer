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

	theta := []float64{1.1796274803348248, 0.45350283575717365, 1.2265378863189635, 0.010213849936610016, 0.019863311713379954, -1.2547548536347255, -0.18976113767438424, -0.07269139545000601, -1.3998332064732484, -0.31271170357784905, -0.1942846440892081, -0.9127803210237302, -0.14664332580363557, -0.06133317524296543, 0.9762368237471596, 0.2565805167734834, 0.07011581024428712, 1.0393800798370414}
	newSample := data.Data
	prediction := predict(theta, newSample)

	return prediction
}
