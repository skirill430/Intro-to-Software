import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
import { HomeComponent } from "./home/home.component";
import { LoginComponent } from "./login/login.component";
import { PageNotFoundComponent } from "./page-not-found/page-not-found.component";
import { CartComponent } from "./cart/cart.component";

// TODO: make if url typed in is not recognized, take user to an error page
// Array of routes (if the path value is typed into URL, take use to specified component)
const routes: Routes = [
    // Default page is search page
    {path: '', component: HomeComponent},
    {path: 'home', component: HomeComponent},
    {path: 'login', component: LoginComponent},
    {path: 'pagenotfound', component: PageNotFoundComponent},
    {path: 'cart', component: CartComponent},
    {path: '**', component: PageNotFoundComponent},

];

@NgModule({
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})

export class AppRoutingModule {}
export const routingComponents = [LoginComponent, HomeComponent]