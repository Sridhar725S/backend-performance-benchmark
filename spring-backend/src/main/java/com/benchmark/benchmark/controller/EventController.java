package com.benchmark.benchmark.controller;

import com.benchmark.benchmark.model.Event;
import com.benchmark.benchmark.service.EventService;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/events")
@RequiredArgsConstructor
public class EventController {

    private final EventService service;

    @PostMapping
    public Event create(
            @RequestBody Event event
    ) {
        return service.create(event);
    }

    @GetMapping("/user/{id}")
    public List<Event> getByUser(
            @PathVariable String id
    ) {
        return service.getByUser(id);
    }

    @GetMapping("/recent")
    public List<Event> getRecent() {
        return service.getRecent();
    }
}