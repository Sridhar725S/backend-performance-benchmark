package com.benchmark.benchmark.repository;

import com.benchmark.benchmark.model.User;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface UserRepository
        extends MongoRepository<User, String> {
}