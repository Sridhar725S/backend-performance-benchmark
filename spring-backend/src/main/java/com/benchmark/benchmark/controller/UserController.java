package com.benchmark.benchmark.controller;

import com.benchmark.benchmark.model.User;
import com.benchmark.benchmark.service.UserService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/users")
@RequiredArgsConstructor
public class UserController {

    private final UserService service;

    @PostMapping
    public User create(
            @RequestBody User user
    ) {
        return service.create(user);
    }

    @GetMapping
    public List<User> getAll() {
        return service.getAll();
    }

    @GetMapping("/{id}")
    public User getById(
            @PathVariable String id
    ) {
        return service.getById(id);
    }
}