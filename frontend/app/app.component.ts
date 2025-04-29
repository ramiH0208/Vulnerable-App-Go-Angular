import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-root',
  template: `
    <h1>Vulnérabilité XSS</h1>
    <form (submit)="submitMessage($event)">
      <input type="text" [(ngModel)]="message" name="message" required>
      <button type="submit">Envoyer</button>
    </form>
    <p [innerHTML]="response"></p>
  `
})
export class AppComponent {
  message = '';
  response = '';

  constructor(private http: HttpClient) {}

  submitMessage(event: Event) {
    event.preventDefault();
    this.http.post<{response: string}>('http://localhost:8080/submit', { message: this.message })
      .subscribe(res => this.response = res.response);
  }
}