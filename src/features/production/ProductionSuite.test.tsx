import { render, screen } from "@testing-library/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { describe, expect, it } from "vitest";
import { ProductionSuite } from "./ProductionSuite";

describe("ProductionSuite", () => {
  it("renders the production console", () => {
    render(
      <QueryClientProvider client={new QueryClient()}>
        <ProductionSuite />
      </QueryClientProvider>
    );

    expect(screen.getByRole("heading", { name: /Musician Production Suite/i })).toBeInTheDocument();
    expect(screen.getByText(/Drop or choose/i)).toBeInTheDocument();
  });
});
