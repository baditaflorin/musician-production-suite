import { z } from "zod";

const artifactSchema = z.object({
  name: z.string(),
  kind: z.string(),
  url: z.string(),
  sizeBytes: z.number().int().nonnegative(),
});

export const jobSchema = z.object({
  id: z.string(),
  filename: z.string(),
  status: z.enum(["queued", "running", "succeeded", "failed"]),
  progress: z.number().min(0).max(100),
  message: z.string(),
  createdAt: z.string(),
  updatedAt: z.string(),
  bpm: z.number().nullable(),
  key: z.string().nullable(),
  chords: z.array(
    z.object({ start: z.number(), end: z.number(), chord: z.string() }),
  ),
  artifacts: z.array(artifactSchema),
  error: z.string().nullable(),
});

export type Job = z.infer<typeof jobSchema>;
export type Artifact = z.infer<typeof artifactSchema>;

const defaultApiBase =
  import.meta.env.VITE_API_BASE_URL ?? "http://localhost:8080";

export function getStoredApiBase(): string {
  return localStorage.getItem("mps.apiBase") ?? defaultApiBase;
}

export function setStoredApiBase(value: string): void {
  localStorage.setItem("mps.apiBase", value.replace(/\/$/, ""));
}

export async function createJob(
  file: File,
  apiBase = getStoredApiBase(),
): Promise<Job> {
  const body = new FormData();
  body.append("audio", file);
  const response = await fetch(`${apiBase}/api/v1/jobs`, {
    method: "POST",
    body,
  });
  return parseJobResponse(response);
}

export async function getJob(
  id: string,
  apiBase = getStoredApiBase(),
): Promise<Job> {
  const response = await fetch(
    `${apiBase}/api/v1/jobs/${encodeURIComponent(id)}`,
  );
  return parseJobResponse(response);
}

async function parseJobResponse(response: Response): Promise<Job> {
  const data: unknown = await response.json();
  if (!response.ok) {
    const message =
      typeof data === "object" && data !== null && "error" in data
        ? String((data as { error: unknown }).error)
        : `Request failed with ${response.status}`;
    throw new Error(message);
  }
  return jobSchema.parse(data);
}
