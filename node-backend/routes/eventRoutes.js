const router = require('express').Router();
const controller = require('../controllers/eventController');

router.post('/', controller.createEvent);
router.get('/user/:id', controller.getEventsByUser);
router.get('/recent', controller.getRecentEvents);

module.exports = router;