"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const mongoose_1 = __importDefault(require("mongoose"));
const projectSchema = new mongoose_1.default.Schema({
    title: {
        type: String,
        required: true,
    },
    description: {
        type: String,
        require: true,
    },
    thumbnail: {
        type: String,
        required: true,
    },
    repoLink: {
        type: String,
        required: true,
    },
    hostLink: {
        type: String,
        required: true,
    },
});
const Project = mongoose_1.default.model("Project", projectSchema);
exports.default = Project;
