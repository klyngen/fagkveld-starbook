import { Injectable } from '@angular/core';
import { merge, Observable, Subject, switchMap } from 'rxjs';
import { HttpService } from './http.service';
import { Star } from './models/star';

@Injectable({
  providedIn: 'root'
})
export class StarService {

  private refreshTrigger = new Subject<number>();

  constructor(private httpClient: HttpService) {

  }

  addStar(comment: string, userId: number) {
    this.httpClient.postStar(comment, userId).subscribe(() => {
      this.refreshTrigger.next(userId);
    })
  }

  starObservable(): Observable<Star[]> {
    return merge(this.httpClient.fetchStars(), this.refreshTrigger.pipe(
      switchMap(() => this.httpClient.fetchStars())))
  }
}
