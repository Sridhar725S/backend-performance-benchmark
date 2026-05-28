const axios = require("axios");

async function seedUsers() {

  for(let i = 1; i <= 5000; i++) {

    try {

      await axios.post(
        "http://localhost:8083/users",
        {
          name: `User ${i}`,
          email: `user${i}@gmail.com`,
          age: 20 + (i % 10)
        }
      );

      console.log(`Inserted User ${i}`);

    } catch(err) {

      console.log("Error:", err.message);
    }
  }
}

seedUsers();