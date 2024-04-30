import { faker } from "@faker-js/faker";

describe("Food Dashboard", () => {
  beforeEach(() => {
    cy.visit("/food_entries");
  });

  context("Mobile View", () => {
    beforeEach(() => {
      cy.viewport("iphone-6");
    });
  });

  context("Desktop View", () => {
    beforeEach(() => {
      cy.viewport(1280, 720);
    });
  });
});
