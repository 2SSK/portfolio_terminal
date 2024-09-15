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
exports.getProjects = exports.createProject = void 0;
const Project_1 = __importDefault(require("../models/Project"));
const multer_1 = __importDefault(require("multer"));
const storage = multer_1.default.diskStorage({
    destination: (req, file, cb) => {
        cb(null, "uploads/");
    },
    filename: (req, file, cb) => {
        cb(null, `${Date.now()}-${file.originalname}`);
    },
});
const upload = (0, multer_1.default)({ storage });
exports.createProject = [
    upload.single("thumbnail"),
    (req, res) => __awaiter(void 0, void 0, void 0, function* () {
        var _a;
        try {
            const { title, description, repoLink, hostLink } = req.body;
            const project = new Project_1.default({
                title,
                description,
                thumbnail: (_a = req.file) === null || _a === void 0 ? void 0 : _a.path,
                repoLink,
                hostLink,
            });
            yield project.save();
            res.status(201).json(project);
        }
        catch (error) {
            res.status(400).json({ error: "Unable to create project" });
        }
    }),
];
const getProjects = (req, res) => __awaiter(void 0, void 0, void 0, function* () {
    const projects = yield Project_1.default.find();
    res.json(projects);
});
exports.getProjects = getProjects;
