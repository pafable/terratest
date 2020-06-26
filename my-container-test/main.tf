provider "docker" {
  host = "unix:///var/run/docker.sock"
}

resource "docker_image" "ghostie" {
    name = "ghost:latest"
}

resource "docker_container" "ghost_container" {
    name = "ghost-cont-01"
    image = "${docker_image.ghostie.latest}"
    ports {
        internal = 2368
        external = 80
    }
    depends_on = [
        docker_image.ghostie,
    ]
}