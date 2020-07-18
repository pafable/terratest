provider "docker" {
#   Windows
  host = "tcp://localhost:2375"

#   MacOS
# host = "unix:///var/run/docker.sock"
}

resource "docker_image" "blog_image" {
    name = terraform.workspace == "prod" ? "ghost:latest" : "wordpress:latest"
}

resource "docker_container" "blog_container" {
    name = "cont-01"
    image = docker_image.blog_image.latest
    ports {
        internal = terraform.workspace == "prod" ? "2368" : "8080"
        external = 80
    }
    depends_on = [
        docker_image.blog_image,
    ]
}

output "docker_image_name" {
    value = docker_image.blog_image.name
}