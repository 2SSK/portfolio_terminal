import express from "express";
import dotenv from "dotenv";
import connectDB from "./config/database";
import projectRoutes from "./routes/projectRoutes";
import experienceRoutes from "./routes/experienceRoutes";
import resumeRoutes from "./routes/resumeRoutes";

dotenv.config();
connectDB();

const app = express();
app.use(express.json());

// Routes
app.use("/api", projectRoutes);
app.use("/api", experienceRoutes);
app.use("/api", resumeRoutes);

// Uploads folder for static files
app.use("/uploads", express.static("uploads"));

const PORT = process.env.PORT || 5000;
app.listen(PORT, () => console.log(`Server running on port ${PORT}`));
