package com.saireddyavs.springconfigclient;


import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.context.config.annotation.RefreshScope;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@SpringBootApplication
public class SpringConfigClientApplication {

	public static void main(String[] args) {
		SpringApplication.run(SpringConfigClientApplication.class, args);
	}
}

@RestController
@RefreshScope
class MessageRestController {

	@Value("${server.message:Unable to connect to config server}")
	private String message;

	@RequestMapping("/server/message")
	String getMsg() {
		return this.message;
	}
}
