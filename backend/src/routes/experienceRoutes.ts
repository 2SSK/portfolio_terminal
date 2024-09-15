import express from "express";
import {
  createExperience,
  getExperiences,
} from "../controllers/experienceController";

const router = express.Router();

router.post("/experiences", createExperience);
router.get("/experiences", getExperiences);

export default router;
