package com.benchmark.benchmark.repository;

import com.benchmark.benchmark.model.Event;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.List;

public interface EventRepository
        extends MongoRepository<Event, String> {

    List<Event> findByUserId(String userId);

    List<Event> findTop100ByOrderByTimestampDesc();
}