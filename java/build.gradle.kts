plugins {
    id("java")
    id("com.google.protobuf") version "0.9.6"
    id("maven-publish")
}

java {
    sourceCompatibility = JavaVersion.VERSION_21
}

dependencies {
    implementation("com.google.protobuf:protobuf-java:4.33.5")
    implementation("io.grpc:grpc-protobuf:1.79.0")
    implementation("io.grpc:grpc-stub:1.79.0")
    implementation("javax.annotation:javax.annotation-api:1.3.2")
}

protobuf {
    protoc { artifact = "com.google.protobuf:protoc:4.33.4" }
    plugins {
        register("grpc") { artifact = "io.grpc:protoc-gen-grpc-java:1.79.0" }
    }
    generateProtoTasks {
        all().configureEach {
            plugins { register("grpc") }
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
