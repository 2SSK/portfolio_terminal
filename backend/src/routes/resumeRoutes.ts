import express from "express";
import { downloadResume } from "../controllers/resumeController";

const router = express.Router();

router.get("/resume", downloadResume);

export default router;
