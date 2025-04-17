import { Component, effect, OnInit } from '@angular/core';
import { HeaderComponent } from '../header/header.component';
import { ChatService } from '../../core/services/chat.service';
import { ChatSkeletonComponent } from '../../shared/components/chat-skeleton/chat-skeleton.component';
import { IChat } from '../../core/interfaces/chat.service.interface';
import { ChatMessageComponent } from '../chat-message/chat-message.component';
import { Router } from '@angular/router';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-chat',
  imports: [HeaderComponent, ChatSkeletonComponent, ChatMessageComponent, FormsModule],
  templateUrl: './chat.component.html',
  styleUrl: './chat.component.scss'
})
export class ChatComponent {
  isLoading: boolean = true;
  codeChats: IChat[] = [];
  CodeBaseId: string = "";
  userQuestion: string = '';
  constructor(
    private chatService: ChatService,
    private router: Router
  ) {
    effect(() => {
      this.CodeBaseId = this.chatService.CodeBaseChatIdentifierSignal();
      this.getChatsForCodeBase();
    })
  }

  private async getChatsForCodeBase() {
    if (this.CodeBaseId) {
      this.codeChats = await this.chatService.RetriveChatHistory(this.CodeBaseId);
      console.log(this.codeChats);
      this.isLoading = false;
      if (!this.codeChats) {
        this.populateDummyChats();
      }
    } else {
      this.router.navigate(['/']);
    }
  }
  appendUserQuestionToChatHistory(question: string) {
    const newChat: IChat = {
      chat_id: "123",
      user_id: "123",
      code_base_id: this.CodeBaseId,
      ai_answer: "",
      user_question: question,
    };
    this.codeChats.push(newChat);
    console.log(this.codeChats);
  }
  updateChat(chat: IChat) {
    if (this.codeChats.length > 0) {
      this.codeChats.pop();
    }
    this.codeChats.push(chat);
  }

  protected async handleSend() {
    if (this.userQuestion.trim()) {
      this.appendUserQuestionToChatHistory(this.userQuestion)
      const chat = await this.chatService.ChatWithCodeBase(this.userQuestion, this.CodeBaseId);
      this.updateChat(chat);
    }
  }
  protected populateDummyChats() {

    this.codeChats = [{
      chat_id: "123",
      user_id: "123",
      code_base_id: this.CodeBaseId,
      ai_answer: "",
      user_question: "Welcome To Briznigr's  Code Assistant . Please Chat On .",
    }]
  }

}
