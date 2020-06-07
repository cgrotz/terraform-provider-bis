

provider "bis" {
  client_id = var.suite_client_id
  client_secret = var.suite_client_secret
}

data "bis_things" "things_instance" {
  solution_id = var.things_solution_id
  api_token = var.things_api_key
}

resource "bis_things_namespace" "my-namespace" {
  solution_id = data.bis_things.things_instance.id
  namespace = "de.cgrotz.test1"
}