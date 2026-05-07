import { beforeEach, describe, expect, it } from "vitest";
import { readRecentJobs, rememberJob } from "./recentJobs";

describe("recent job storage", () => {
  beforeEach(() => localStorage.clear());

  it("keeps newest jobs first and de-duplicates ids", () => {
    rememberJob("a");
    rememberJob("b");
    rememberJob("a");

    expect(readRecentJobs()).toEqual(["a", "b"]);
  });
});

