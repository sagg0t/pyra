import htmx from "htmx.org/dist/htmx.esm";
import { Application } from "@hotwired/stimulus";
import { ProductSearchController } from "./controllers/product_search_controller";
import { ProductRowController } from "./controllers/product_row_controller";

window.htmx = htmx;
// htmx.logAll();

const stimulusApp = Application.start();
stimulusApp.debug = true;
window.C = stimulusApp;

stimulusApp.register("product-search", ProductSearchController);
stimulusApp.register("product-row", ProductRowController);
