const key = "mps.recentJobs";

export function readRecentJobs(): string[] {
  const raw = localStorage.getItem(key);
  if (!raw) return [];
  try {
    const parsed = JSON.parse(raw);
    return Array.isArray(parsed)
      ? parsed.filter((item) => typeof item === "string")
      : [];
  } catch {
    return [];
  }
}

export function rememberJob(id: string): string[] {
  const jobs = [
    id,
    ...readRecentJobs().filter((existing) => existing !== id),
  ].slice(0, 8);
  localStorage.setItem(key, JSON.stringify(jobs));
  return jobs;
}
