import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AppComponent } from '../app.component';
import { HttpService } from '../http.service';
import { MatSidenavModule } from '@angular/material/sidenav';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatSortModule } from '@angular/material/sort';
import { FormsModule } from '@angular/forms';
import {HttpClientModule} from '@angular/common/http';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule, routingComponents } from '../app-routing.module';
import { HomeComponent } from './home.component';
import {MatTableModule} from '@angular/material/table';
import { MatTableDataSource } from '@angular/material/table';


describe('HomeComponent', () => {
  let component: HomeComponent;
  let fixture: ComponentFixture<HomeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, MatSidenavModule, MatSlideToggleModule, MatToolbarModule, MatIconModule, MatButtonModule, FormsModule,
      MatFormFieldModule, MatInputModule, BrowserAnimationsModule, ReactiveFormsModule, AppRoutingModule, MatTableModule, MatSortModule],
      providers: [HttpService],
      declarations: [ HomeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(HomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  
  it('should render a search text box', () => {
    const fixture = TestBed.createComponent(HomeComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector("mat-form-field input[name='Input']")).toBeTruthy();
  });

  it('should render a toolbar', () => {
    const fixture = TestBed.createComponent(HomeComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector("mat-drawer-container mat-toolbar[color='primary']")).toBeTruthy();
  });

  it('should render a sidebar', () => {
    const fixture = TestBed.createComponent(HomeComponent);
    fixture.detectChanges();
    const compiled = fixture.debugElement.nativeElement;
    expect(compiled.querySelector("mat-drawer-container mat-drawer[mode='side']")).toBeTruthy();
  });
});
