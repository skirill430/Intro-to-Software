import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  title = 'Login Page';
  constructor(private router:Router){}

  // // function to take user to home page (not used)
  // goToPage(pageName:string):void {
  //   this.router.navigate([`${pageName}`]);
  // }
}
