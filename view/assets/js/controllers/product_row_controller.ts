import { Controller } from "@hotwired/stimulus";
import type { ProductSelectEvent } from "./product_search_controller";
import type { Product } from "../types/product";
import { assert } from "../assert";

export class ProductRowController extends Controller {
    static values = {
        isLast: Boolean
    };

    declare isLastValue: boolean;

    static targets = [
        "amountInput",
        "calories",
        "proteins",
        "fats",
        "carbs"
    ];

    declare readonly amountInputTarget: HTMLInputElement;
    declare readonly caloriesTarget: HTMLElement;
    declare readonly proteinsTarget: HTMLElement;
    declare readonly fatsTarget: HTMLElement;
    declare readonly carbsTarget: HTMLElement;

    productSelect(e: ProductSelectEvent) {
        const product: Product = e.detail;

        this.amountInputTarget.value = (100).toString();
        this.caloriesTarget.innerText = product.calories.toString();
        this.proteinsTarget.innerText = product.proteins.toString();
        this.fatsTarget.innerText = product.fats.toString();
        this.carbsTarget.innerText = product.carbs.toString();

        if (this.isLastValue) {
            this.isLastValue = false;
            this.appendNextRow();
        }
    }

    appendNextRow() {
        const template = document.getElementById("product-row-template") as HTMLTemplateElement | null;
        assert(!!template, "missing product row template");

        const newRow = template.content.cloneNode(true);
        this.element.parentElement?.appendChild(newRow);
    }

    destroy() {
        if (this.isLastValue) {
            return;
        }

        this.element.remove();
    }
};
