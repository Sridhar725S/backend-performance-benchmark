require('dotenv').config();

const express = require('express');
const cors = require('cors');

const connectDB = require('./config/db');

const app = express();

connectDB();

app.use(cors());
app.use(express.json());

app.use('/users', require('./routes/userRoutes'));
app.use('/events', require('./routes/eventRoutes'));

app.listen(process.env.PORT, () => {
  console.log(`Server running on ${process.env.PORT}`);
});