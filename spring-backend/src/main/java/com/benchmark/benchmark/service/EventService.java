package com.benchmark.benchmark.service;

import com.benchmark.benchmark.model.Event;
import com.benchmark.benchmark.repository.EventRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class EventService {

    private final EventRepository repository;

    public Event create(Event event) {
        return repository.save(event);
    }

    public List<Event> getByUser(String userId) {
        return repository.findByUserId(userId);
    }

    public List<Event> getRecent() {
        return repository
                .findTop100ByOrderByTimestampDesc();
    }
}