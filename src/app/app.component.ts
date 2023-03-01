import { Component } from '@angular/core';
// -- Code that was moved into 'home' folder (saving in case for future) --
// import { HttpService } from './http.service';
// import * as http from './http.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']

})
export class AppComponent {
  title = 'Quick Shop';

  // -- old moved code cont. --
  // items : http.RootObject;
  // search : string;
  // showFiller = false;

  // constructor(private httpService: HttpService) {}

  // save() {
  //   console.log(this.search);
  //   this.httpService.getWalmartItems(this.search).subscribe(response => {
  //     console.log(response);
  //     this.items = response;
  //   })
  // }
}
