const Event = require('../models/Event');

exports.createEvent = async (req, res) => {
  const event = await Event.create(req.body);
  res.json(event);
};

exports.getEventsByUser = async (req, res) => {
  const events = await Event.find({
    userId: req.params.id
  });

  res.json(events);
};

exports.getRecentEvents = async (req, res) => {
  const events = await Event.find()
    .sort({ timestamp: -1 })
    .limit(100);

  res.json(events);
};