"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const dotenv_1 = __importDefault(require("dotenv"));
const database_1 = __importDefault(require("./config/database"));
const projectRoutes_1 = __importDefault(require("./routes/projectRoutes"));
const experienceRoutes_1 = __importDefault(require("./routes/experienceRoutes"));
const resumeRoutes_1 = __importDefault(require("./routes/resumeRoutes"));
dotenv_1.default.config();
(0, database_1.default)();
const app = (0, express_1.default)();
app.use(express_1.default.json());
// Routes
app.use("/api", projectRoutes_1.default);
app.use("/api", experienceRoutes_1.default);
app.use("/api", resumeRoutes_1.default);
// Uploads folder for static files
app.use("/uploads", express_1.default.static("uploads"));
const PORT = process.env.PORT || 5000;
app.listen(PORT, () => console.log(`Server running on port ${PORT}`));
