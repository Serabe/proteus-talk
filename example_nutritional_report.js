const messages = require('./generated_pb');
const services = require('./generated_grpc_pb');
const grpc = require('grpc');

let client = new services.GomadServiceClient('localhost:8182', grpc.credentials.createInsecure());

let pizza = new messages.Pizza();

console.log("Pizza ingredients: ", pizza.getIngredientsList().length);

let request = new messages.AddIngredientRequest();
request.setArg1(pizza);
request.setArg2(messages.Ingredient.PEPPERONI);
request.setArg3(3);

client.addIngredient(request, function(err, response) {
  if(err) {
    console.log("FAILED ", err);
  }

  console.log(response.getIngredientsList()[0].getIngredient() == messages.Ingredient.PEPPERONI);
  console.log(response.getIngredientsList()[0].getPortions());

  pizza = response;

  client.nutritionalReport(pizza, function(err, response) {
    if (err) {
      console.log('Something bad happened', err);
    }

    console.log(response.getResult1());
  });
});
