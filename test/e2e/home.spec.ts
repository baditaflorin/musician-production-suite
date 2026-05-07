import { expect, test } from "@playwright/test";

test("production console loads", async ({ page }) => {
  await page.goto("/");
  await expect(
    page.getByRole("heading", { name: "Musician Production Suite" }),
  ).toBeVisible();
  await expect(page.getByText(/New production job/i)).toBeVisible();
});
