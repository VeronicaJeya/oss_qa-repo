package com.thirddemo3;


import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.GetMapping;

@RestController
@RequestMapping("/first")

public class Thirdcontroller3 {
	@GetMapping
	public String today() {
		String aString = java.time.LocalDate.now().toString(); 
		String bString = java.time.LocalTime.now().toString(); 
		String cString = aString + bString;
		return cString;
	}

}


