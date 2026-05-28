package com.benchmark.benchmark.model;

import lombok.Data;
import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import java.time.LocalDateTime;

@Data
@Document(collection = "events")
public class Event {

    @Id
    private String id;

    private String userId;

    private String type;

    private LocalDateTime timestamp =
            LocalDateTime.now();
}