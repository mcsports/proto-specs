plugins {
    id("java")
    alias(libs.plugins.protobuf)
    id("maven-publish")
}

java {
    sourceCompatibility = JavaVersion.VERSION_21
}

dependencies {
    implementation(libs.protobuf.java)
    implementation(libs.grpc.protobuf)
    implementation(libs.grpc.stub)
    implementation(libs.javax.annotations)
}

protobuf {
    protoc { artifact = libs.protobuf.compiler.get().toString() }
    plugins {
        register("grpc") { artifact = libs.grpc.java.generator.get().toString() }
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
