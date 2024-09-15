import express from "express";
import { createProject, getProjects } from "../controllers/projectController";

const router = express.Router();

router.post("/projects", createProject);
router.get("/projects", getProjects);

export default router;
