"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const experienceController_1 = require("../controllers/experienceController");
const router = express_1.default.Router();
router.post("/experiences", experienceController_1.createExperience);
router.get("/experiences", experienceController_1.getExperiences);
exports.default = router;
