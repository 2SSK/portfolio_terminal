"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.downloadResume = void 0;
const path_1 = __importDefault(require("path"));
const downloadResume = (req, res) => {
    const filePath = path_1.default.join(__dirname, "../../uploads/resume.pdf");
    res.download(filePath);
};
exports.downloadResume = downloadResume;
