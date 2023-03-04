const { createApp } = Vue

  createApp({
    data() {
      return {
        message: getPredictedData("1-4d22f711-afae-474d-8443-9ca03943f843").then(data => console.log(data))
      }
    }
  }).mount('#app')