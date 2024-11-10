import { Controller } from "@hotwired/stimulus";
import type { Product } from "../types/product";
import { throttle } from "@github/mini-throttle";
import { assert } from "../assert";

export type ProductSelectEvent = CustomEvent<Product>;
export type ProductClearEvent = CustomEvent;

const product: Product = {
    ID: "asdf",
    calories: 100,
    proteins: 5,
    fats: 2,
    carbs: 3
};

export class ProductSearchController extends Controller {
    static values = {
        "datalistId": String
    };
    declare datalistIdValue: string;

    searchProducts = throttle(async (e: InputEvent): void => {
        const input = e.target as HTMLInputElement;
        const datalist = this.findDatalistEl();

        if (input.value === "") {
            datalist.innerHTML = "";
            return;
        }

        const params = new URLSearchParams({ q: input.value });

        const path = "/foodProducts/search";
        const fullURI = path + "?" + params.toString();

        const res = await fetch(fullURI, { method: "POST" });
        const products = await res.json();

        const datalistOptions = products.map(x => {
            const option = document.createElement("option");
            option.setAttribute("value", x.Name);

            return option;
        })

        datalist.replaceChildren(...datalistOptions);
    }, 500, { start: false });

    searchSubmit(e: Event): void {
        const input = e.target as HTMLInputElement;
        if (input.value === "") {
            const clearEvent: ProductClearEvent = new CustomEvent("product-clear");
            this.element.dispatchEvent(clearEvent)

            this.findDatalistEl().innerHTML = "";
        } else {
            const selectEvent: ProductSelectEvent = new CustomEvent("product-select", { detail: product });
            this.element.dispatchEvent(selectEvent)
        }
    }

    findDatalistEl(): HTMLElement {
        const datalist = document.getElementById(this.datalistIdValue);
        assert(!!datalist, "datalist element for the search result is missing");

        return datalist;
    }
};
