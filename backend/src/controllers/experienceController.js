"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getExperiences = exports.createExperience = void 0;
const Experience_1 = __importDefault(require("../models/Experience"));
const createExperience = (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const experience = new Experience_1.default(req.body);
        yield experience.save();
        res.status(201).json(experience);
    }
    catch (error) {
        res.status(400).json({ error: "Unable to create experience" });
    }
});
exports.createExperience = createExperience;
const getExperiences = (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const experiences = yield Experience_1.default.find();
    res.json(experiences);
});
exports.getExperiences = getExperiences;
