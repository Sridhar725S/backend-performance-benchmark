const router = require('express').Router();
const controller = require('../controllers/userController');

router.post('/', controller.createUser);
router.get('/', controller.getUsers);
router.get('/:id', controller.getUserById);

module.exports = router;