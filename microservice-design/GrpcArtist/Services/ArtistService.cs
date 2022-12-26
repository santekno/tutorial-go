using Grpc.Core;
using GrpcArtist;

namespace GrpcArtist.Services;

public class ArtistService : Artist.ArtistBase
{
  List<ArtistResponse> artists = new List<ArtistResponse>();
  private readonly ILogger<ArtistService> _logger;
  public ArtistService(ILogger<ArtistService> logger)
  {
    _logger = logger;

    artists.Add(new ArtistResponse() { Id = 1, Name = "Peterpan", Rating = 5 });
    artists.Add(new ArtistResponse() { Id = 2, Name = "Nidji", Rating = 4 });
    artists.Add(new ArtistResponse() { Id = 3, Name = "DMasiv", Rating = 4 });
    artists.Add(new ArtistResponse() { Id = 4, Name = "Ungu", Rating = 5 });
  }

  public override Task<ArtistResponse> GetById(ArtistRequest request, ServerCallContext context)
  {
    var artist = artists.Find(x => x.Id.Equals(request.Id));
    if (artist != null)
    {
      return Task.FromResult(new ArtistResponse { Id = artist.Id, Name = artist.Name, Rating = artist.Rating });
    }
    return Task.FromResult(new ArtistResponse { });
  }

  public override Task<ArtistsResponse> GetAll(PagingRequest request, ServerCallContext context)
  {
    var artistPaging = artists.Skip(request.Offset - 1 * request.Limit).Take(request.Limit).ToList();
    ArtistsResponse responses = new ArtistsResponse();
    responses.Artists.AddRange(artistPaging);
    return Task.FromResult(responses);
  }
}
