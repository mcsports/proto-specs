plugins {
    alias(libs.plugins.kotlin)
    alias(libs.plugins.protobuf)
    id("maven-publish")
}

dependencies {
    api(project(":proto-java"))
    implementation(libs.protobuf.kotlin)

    implementation(libs.grpc.protobuf)
    implementation(libs.grpc.stub)
    implementation(libs.grpc.kotlin.stub)
}

protobuf {
    protoc { artifact = libs.protobuf.compiler.get().toString() }
    plugins {
        register("grpc") {
            artifact = libs.grpc.java.generator.get().toString()
        }
        register("grpckt") { artifact = libs.grpc.kotlin.generator.get().toString() }
    }
    generateProtoTasks {
        all().configureEach {
            plugins {
                register("grpc")
                register("grpckt")
            }
        }
    }
}

publishing {
    repositories {
        val reposiliteUrl = findProperty("reposiliteUrl") as String?
        if(reposiliteUrl != null) {
            maven {
                name = "Reposilite"
                url = uri(reposiliteUrl)
                credentials {
                    username = findProperty("reposiliteUsername") as String?
                    password = findProperty("reposilitePassword") as String?
                }
            }
        }
    }

    publications {
        create<MavenPublication>("maven") {
            from(components["java"])
        }
    }
}