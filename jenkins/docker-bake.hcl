variable "VERSION" {
  default = "2.440.3-lts-jdk17"
}

variable "FIXID" {
  default = "1"
}

group "default" {
  targets = ["jenkins-amd64", "jenkins-arm64"]
}

target "jenkins-amd64" {
    labels = {
        "cloud.opsbox.author" = "seanly"
        "cloud.opsbox.image.name" = "jenkins"
        "cloud.opsbox.image.version" = "${VERSION}"
        "cloud.opsbox.image.fixid" = "${FIXID}"
    }
    dockerfile = "Dockerfile"
    context  = "./"
    args = {
        VERSION="${VERSION}"
    }
    platforms = ["linux/amd64"]
    tags = [
        "registry.cn-chengdu.aliyuncs.com/seanly/test:jenkins-${VERSION}-${FIXID}",
        "registry.cn-hangzhou.aliyuncs.com/seanly/test:jenkins-${VERSION}-${FIXID}"
    ]
    output = ["type=image,push=true"]
}

target "jenkins-arm64" {
    labels = {
        "cloud.opsbox.author" = "seanly"
        "cloud.opsbox.image.name" = "jenkins"
        "cloud.opsbox.image.version" = "${VERSION}"
        "cloud.opsbox.image.fixid" = "${FIXID}"
    }
    dockerfile = "Dockerfile"
    context  = "./"
    args = {
        VERSION="${VERSION}"
    }
    platforms = ["linux/arm64"]
    tags = [
        "registry.cn-chengdu.aliyuncs.com/seanly/test:jenkins-${VERSION}-${FIXID}-arm64",
        "registry.cn-hangzhou.aliyuncs.com/seanly/test:jenkins-${VERSION}-${FIXID}-arm64"
    ]
    output = ["type=image,push=true"]
}
