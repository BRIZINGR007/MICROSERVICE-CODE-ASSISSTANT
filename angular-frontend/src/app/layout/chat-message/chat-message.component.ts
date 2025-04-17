import { Component, Input } from '@angular/core';
import { IChat } from '../../core/interfaces/chat.service.interface';
import { ChatSkeletonComponent } from '../../shared/components/chat-skeleton/chat-skeleton.component';

@Component({
  selector: 'app-chat-message',
  imports: [ChatSkeletonComponent],
  templateUrl: './chat-message.component.html',
  styleUrl: './chat-message.component.scss'
})
export class ChatMessageComponent {
  @Input() chat!: IChat;

}
